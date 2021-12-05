WITH temp AS(
    select b.id as id, count(bs.seat_number) as booked_seats_count
    from bus b
    left join bus_seat bs on b.id = bs.bus_id
    group by b.id
    )
    select b.id, b.start_time, b.end_time, b.name, b.seat_count, 
    d.from_place || '-' || d.to_place AS destination_name, 
    is_full, b.seat_count - t.booked_seats_count as remaining_seats
    from temp t
    join bus b on t.id = b.id
    JOIN destination d ON b.destination_id = d.id
    WHERE b.destination_id = $1;




        SELECT 
        b.id, b.start_time, b.end_time, b.name, b.seat_count,  
        d.from_place || '-' || d.to_place AS destination_name,
        count(bs.seat_number)
        FROM bus b
        JOIN destination d ON b.destination_id = d.id
        JOIN bus_seat bs ON bs.bus_id = b.id
        WHERE b.destination_id = 4
        ORDER BY b.created_at desc