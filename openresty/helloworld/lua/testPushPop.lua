local dogs = ngx.shared.dogs

local len, err = dogs:lpush("foo", "bar")
if len then
	ngx.say("push success")
else
	ngx.say("push err: ", err)
end

local val, err = dogs:llen("foo")
ngx.say(val, " ", err)


local val, err = dogs:lpop("foo")
ngx.say(val, " ", err)

local val, err = dogs:llen("foo")
ngx.say(val, " ", err)


local val, err = dogs:lpop("foo")
ngx.say(val, " ", err)
