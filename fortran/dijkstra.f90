!a wrapper for the example with Dijkstra's algorithm
subroutine exampleDijkstra(G,start,end)
    use iso_fortran_env, wp => real64
    use graphf
    implicit none
    type(graph) :: G
    integer :: start,end
    real(wp),allocatable :: dist(:)
    integer,allocatable :: prev(:)
    integer,allocatable :: path(:)
    real(wp) :: dummy
    integer :: lpath
    integer :: i

    !-- allocate space for the dist and prev vectors
    allocate(dist(G%V),prev(G%V))

    !-- run Dijkstra's algo, get distance and prev
    call Dijkstra(G, start, dist, prev)
 
    !-- allocate an array for analyzing the path
    allocate(path(G%V), source = 0)
    if(dist(end) == huge(dummy))then
         write(*,'(a,i0,a,i0)') 'There is no path from vertex ',start,' to vertex ',end
    else
        lpath = 0
        call getPathD(G%V,lpath,start,end,prev,path)
        write(*,'(a,i0,a,i0,a)') "shortest path from vertex ", start, " to vertex ", end, ":"
        do i=1,lpath
           write(*,'(1x,i0)',advance='no') path(i)
        enddo
        write(*,*)
        write(*,'(a,f12.4)') 'with a total path length of ',dist(end)
    endif

    deallocate(path,prev,dist)
    return
end subroutine

subroutine Dijkstra(G, start, dist, prev)
    use iso_fortran_env, wp => real64
    use graphf
    implicit none
    type(graph) :: G
    integer,intent(in) :: start
    real(wp),intent(inout) :: dist(*)
    integer,intent(inout) :: prev(*)
    real(wp) :: inf
    real(wp) :: newdist
    integer,allocatable :: Q(:)
    integer :: Qcount
    integer :: minQ !this is a function
    integer :: i,j,u,v

    inf = huge(inf)
    !-- allocate array of unvisited vertices,
    !   distances and previously visited nodes
    allocate(Q(G%V))
    Qcount=G%V
    do i=1,G%V
        Q(i)=i
        prev(i) = -1
        dist(i) = inf
    enddo
    prev(start) = start
    dist(start) = 0.0_wp

    !-- as long as there are unvisited vertices in Q
    do while (Qcount >= 1)
        !-- get vertix with smallest reference value in dist
        j = minQ(Q,dist,Qcount)
        v = Q(j)            !track the vertex k
        Q(j) = Q(Qcount)    !overwrite Q(j)
        Qcount = Qcount - 1 !"resize" Q
        !-- loop over the neighbours of vertex k
        do u=1,G%V
            if(G%nmat(v,u)==0) cycle
            if(prev(u).ne.v)then
            newdist = dist(v) + G%emat(v,u)
            !-- is the new dist value of u better than the previous one?
            if( newdist < dist(u) )then
                dist(u) = newdist
                prev(u) = v
            endif
            endif
        enddo
    enddo
    deallocate(Q)
    return
end subroutine

function minQ(Q,dist,Qcount)
    use iso_fortran_env, only: wp => real64, error_unit
    implicit none
    integer :: minQ
    integer :: Q(*)
    real(wp) :: dist(*)
    real(wp) :: dref,d
    integer :: Qcount
    integer :: i,j

    dref = huge(dref)
    do i=1,Qcount
        j = Q(i)
        d = dist(j)
        if(d < dref)then
            minQ = i !the position within Q has to be returned, not the vertex!
            dref = d
        endif
    enddo
    if(dref == huge(d))then
        write(error_unit,*) "warning: disconnected graph!"
    endif
    return
end function

recursive subroutine getPathD(nmax, lpath, start, pos, prev, path)
    implicit none
    integer :: nmax
    integer :: lpath
    integer :: start
    integer :: pos
    integer :: prev(nmax)
    integer :: path(nmax)
    integer :: i,j,k
    lpath = lpath + 1
    if(pos == start)then !exit recursion
       path(lpath)  = pos
       do i=1,lpath/2
          j=lpath - i + 1
          k=path(i)
          path(i) = path(j)
          path(j) = k
       enddo
    else
      path(lpath) = pos
      j = prev(pos)
      call getPathD(nmax, lpath, start, j, prev, path)
    endif
    return
end subroutine getPathD
