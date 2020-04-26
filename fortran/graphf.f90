! A simple fortran module for setting up a graph.
! Just vertices, edges and neighbours. 
module graphf
    use iso_fortran_env, wp => real64, sp => real32
    implicit none
    private
    public :: graph
    public :: example1

    type graph
       integer :: V = 0
       integer :: E = 0
       integer,allocatable  :: edges(:,:)
       real(wp),allocatable :: eweight(:) 
       real(sp),allocatable :: emat(:,:)
       integer,allocatable :: nmat(:,:)   !matrix of neighbours (1 or 0)
       logical :: directed = .false.
       
       contains
        procedure :: deallocate => deallocate_graph
        procedure :: addVertex => addVertex_graph
        procedure :: setOrder => addVertex_graph
        procedure :: addEdge => addEdge_graph
        procedure :: degree => degree_graph

    end type graph

contains

subroutine deallocate_graph(self)
    implicit none
    class(graph) :: self
    if(allocated(self%edges)) deallocate(self%edges)
    if(allocated(self%eweight)) deallocate(self%eweight)
    if(allocated(self%nmat)) deallocate(self%nmat)
    if(allocated(self%emat)) deallocate(self%emat)
    self%V = 0
    self%E = 0
    return
end subroutine deallocate_graph

!-- add a vertex (or x vertices) to the graph
subroutine addVertex_graph(self,x)
    class(graph) :: self
    integer,optional :: x !number of new vertices
    integer,allocatable :: ndum(:,:)
    real(sp),allocatable :: edum(:,:)
    integer :: k,o,n
    if(present(x))then
        if(x<1)return
        k = x
    else
        k = 1
    endif
   
    self%V = self%V + k
    if(.not.allocated(self%nmat))then
        allocate(self%nmat(k,k), source = 0)
        allocate(self%emat(k,k), source = 0.0_sp)
    else
        o = self%V - k
        n = self%V
        allocate(ndum(n,n),edum(n,n))
        ndum(1:o,1:o) = self%nmat(1:o,1:o)
        edum(1:o,1:o) = self%emat(1:o,1:o)
        call move_alloc(ndum,self%nmat)
        call move_alloc(edum,self%emat)
        self%nmat(:,n) = 0
        self%nmat(n,:) = 0
        self%emat(:,n) = 0
        self%emat(n,:) = 0
    endif
    
end subroutine addVertex_graph

!-- add an edge to the graph
subroutine addEdge_graph(self, vertex1, vertex2, weight)
    class(graph) :: self
    integer,intent(in) :: vertex1,vertex2
    real(sp),optional :: weight
    real(sp) :: w2
    integer :: edge(2)
    !integer,allocatable :: edum(:,:)
    !real(wp),allocatable :: wdum(:)
    logical :: ex
    !-- if any of the two vertices is larger than the specified number
    !   of vertices in the graph, print a warning and stop
    if(vertex1 > self%V .or. vertex2 > self%V)then
        error stop "please set the number of vertices before defining edges!"
    endif
    !-- the edge weight is an optional argument. If not present it is = 1.0
    if(present(weight))then
        w2=weight
    else
        w2=1.0_sp
    endif
    !-- only undirected graphs are sorted
    if(vertex1 > vertex2 .and. .not.self%directed)then
        edge(2) = vertex1
        edge(1) = vertex2
    else       
        edge(1) = vertex1
        edge(2) = vertex2
    endif 
    !-- check if the vertex was already documented in the neighbour matrix
    ex=areNeighbours(self,edge(1),edge(2))
    if(.not.ex)then
    self%nmat(edge(1),edge(2)) = 1
    self%emat(edge(1),edge(2)) = w2
    if(.not.self%directed)then !for undirected graphs nmat is symmetric
        self%nmat(edge(2),edge(1)) = 1
        self%emat(edge(2),edge(1)) = w2
    endif    
    self%E = self%E + 1
    endif

    return
end subroutine addEdge_graph

!-- are two edges identical?
logical function equalEdge(edge1, edge2)
    implicit none
    integer,intent(in) :: edge1(2)
    integer,intent(in) :: edge2(2)
    equalEdge = .false.
    if( edge1(1)==edge2(1) .and. edge1(2)==edge2(2))then
        equalEdge = .true.
    else
        equalEdge = .false.
    endif
    return
end function equalEdge

!-- decide if the two vertices are neighbouring
logical function areNeighbours(self,v1,v2)
    implicit none
    class(graph) :: self
    integer,intent(in) :: v1,v2
    areNeighbours = .false.
    if( self%nmat(v1,v2) == 1)then
        areNeighbours = .true.
    endif
    return
end function areNeighbours

!-- get the degree of a vertex (i.e., the number of its neighbours)
function degree_graph(self, v, neighbours)
    class(graph),intent(in) :: self
    integer :: degree_graph
    integer,intent(in) :: v
    integer :: i,u
    integer :: neighbours(*)
    u = 0
    do i=1,self%V
        !-- since nmat is *not* symmetric for directed graphs
        !   this should work for both, directed and undirected
        if( self%nmat(v,i)==1)then
            u = u + 1
            neighbours(u)=i
        endif
    enddo
    degree_graph = u
    return
end function degree_graph

!=====================================================================================!

subroutine example1(G)
     implicit none
     class(graph) :: G
     real(sp) :: l
    call G%setOrder(16)
    l = 1.0_sp  

    call G%addEdge( 1, 2,l)
    call G%addEdge( 1, 5, l)
    call G%addEdge( 2, 3, l)
    call G%addEdge( 3, 4, l)
    call G%addEdge( 3, 8, l)
    call G%addEdge( 4, 5, l)
    call G%addEdge( 4, 6, l)
    call G%addEdge( 5, 9, l)
    call G%addEdge( 6, 7, l)
    call G%addEdge( 6, 10, l)
    call G%addEdge( 6, 11, l)
    call G%addEdge( 7, 8, l)
    call G%addEdge( 7, 15, l)
    call G%addEdge( 9, 10, l)
    call G%addEdge(11, 12, l)
    call G%addEdge(11, 13, l)
    call G%addEdge(12, 13, l)
    call G%addEdge(12, 14, l)
    call G%addEdge(13, 15, l)
    call G%addEdge(14, 15, l)
    call G%addEdge(15, 16, l)
    return
end subroutine example1


end module graphf