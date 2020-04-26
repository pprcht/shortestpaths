!a wrapper for the example with the Floyd-Warshall algorithm
subroutine exampleFloydWarshall(G,start,end)
    use iso_fortran_env, wp => real64
    use graphf
    implicit none
    type(graph) :: G
    integer :: start,end
    real(wp),allocatable :: dist(:,:)
    integer,allocatable :: prev(:,:)
    integer,allocatable :: path(:)
    real(wp) :: dummy
    integer :: lpath
    integer :: i

    !-- allocate space for the dist and prev vectors
    allocate(dist(G%V,G%V),prev(G%V,G%V))

    !-- run the Floy-Warshall algo, get distance and prev
    call FloydWarshall(G, dist, prev)
 
    !-- allocate an array for analyzing the path
    allocate(path(G%V), source = 0)
    if(dist(start,end) == huge(dummy))then
         write(*,'(a,i0,a,i0)') 'There is no path from vertex ',start,' to vertex ',end
    else
        call getPathFW(G%V,prev,start,end,path,lpath)
        write(*,'(a,i0,a,i0,a)') "shortest path from vertex ", start, " to vertex ", end, ":"
        do i=1,lpath
           write(*,'(1x,i0)',advance='no') path(i)
        enddo
        write(*,*)
        write(*,'(a,f12.4)') 'with a total path length of ',dist(start,end)
    endif

    deallocate(path,prev,dist)
    return
end subroutine

!-- implementation of the algorithm including setup of "dist" and "prev"
subroutine FloydWarshall(G, dist, prev)
    use iso_fortran_env, wp => real64
    use graphf
    implicit none
    type(graph) :: G
    real(wp),intent(inout) :: dist(G%V,G%V)
    integer,intent(inout) :: prev(G%V,G%V)
    real(wp) :: inf
    real(wp) :: kdist
    integer :: i,j,k

    inf = huge(inf)
    !-- set distances and previously visited nodes
    dist=inf
    prev=-1
    do i=1,G%V
        do j=1,G%V
            if(G%nmat(i,j)==1)then
               dist(i,j) = G%emat(i,j)
               prev(i,j) = j
            endif
        enddo
    enddo

    !-- The algorithm is based on the following assumption:
	!   If a shortest path from vertex u to vertex v runns through a thrid
	!   vertex w, then the paths u-to-w and w-to-v are already minimal.
	!   Hence, the shorest paths are constructed by searching all path
    !   that run over an additional intermediate point k
    !   The following loop is the actual algorithm
    do k=1,G%V
        do i=1,G%V
            do j=1,G%V
                kdist = dist(i,k) + dist(k,j)
                !-- if the path ij runs over k, update
                if(dist(i,j) > kdist)then
                    dist(i,j) = kdist
                    prev(i,j) = prev(i,k)
                endif
            enddo
        enddo
    enddo
    
    return
end subroutine

!-- reconstruct the path start to end for the Floyd-Warshall algorithm
subroutine getPathFW(nmax,prev,start,end,path,lpath)
    implicit none
    integer :: nmax
    integer :: start
    integer :: end
    integer :: prev(nmax,nmax)
    integer :: path(nmax)
    integer :: lpath
    integer :: i,k
    if(prev(start,end) == -1)then
        path(1)=-1
        return
    endif
    k=1
    path(k) = start
    i=start
    do while (  i .ne. end  )
        i = prev(i,end)
        k=k+1
        path(k) = i
    enddo
    lpath=k
    return
end subroutine getPathFW
