local f = io.open("./puzzle1/input.txt", "r")

if (f == nil) then
  print("nil file")
  return
end

local last = 0
local Count = -1
for line in f:lines() do
  local current = tonumber(line)
  if (current == nil) then
    return
  elseif (current > last)
  then
    Count = Count + 1
  end
  last = current
end

print(Count)
