def func(i)
    i > 2 ? func(i-1) + func(i-2) : 1
end

i = 35
s = 10

threads = []
s.times {
    threads.push(Thread.new { p func(i) })
}
threads.each {|t| t.join}
