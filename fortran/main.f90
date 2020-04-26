program MAIN
    use iso_fortran_env, wp => real64
    use graphf
    implicit none
    type(graph) :: G
    integer :: start,end

    !-- set up an example
    call example1(G)
    start = 1 !remember that arrays start at 1 in fortran
    end = 14

    !-- call the example with Dijkstra's Algorithm
    write(*,'(a,i0,a,i0,a)') "Shortest path from vertex ", start, " to vertex ", end, &
    & " using Dijkstra's algorithm:"
    call exampleDijkstra(G,start,end)

    !-- call the example with the Floyd-Warshall Algorithm
    write(*,'(/,a,i0,a,i0,a)') "Shortest path from vertex ", start, " to vertex ", end, &
    & " using the Floyd-Warshall algorithm:"
    call exampleFloydWarshall(G,start,end)

end program MAIN

