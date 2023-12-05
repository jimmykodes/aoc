local f = io.open("assets/input.txt", "r")

if (f == nil) then
  print("nil file")
  return
end

local last = 0
local count = -1
for line in f:lines() do
  local current = tonumber(line)
  if (current == nil) then
    return
  elseif (current > last)
  then
    count = count + 1
  end
  last = current
end

print(count)
