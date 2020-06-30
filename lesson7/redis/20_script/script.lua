if redis.call("GET", KEYS[1]) ~= false then
    return redis.call("INCRBY", KEYS[1], ARGV[1])
end
return false
