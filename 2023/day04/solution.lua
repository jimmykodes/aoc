Card = {}

function Card:new(o)
  o = o or {}
  setmetatable(o, self)
  self.__index = self
  return o
end

function split(str, sep)
  local t = {}
  for s in string.gmatch(str, "([^" .. sep .. "]+)") do
    table.insert(t, s)
  end
  return t
end

function ints(str)
  local out = {}
  for s in string.gmatch(str, "[%d]+") do
    out[#out + 1] = tonumber(s)
  end
  return out
end

function contains(l, val)
  for _, value in ipairs(l) do
    if val == value then
      return true
    end
  end
  return false
end

function load(fn)
  local f = io.open(fn, "r")
  if f == nil then return end
  local out = {}
  for line in f:lines() do
    local spl = split(line, ":")
    local card, nums = spl[1], spl[2]

    spl = split(nums, "|")
    local winning, have = spl[1], spl[2]

    winning = ints(winning)
    have = ints(have)

    local wins = 0
    for _, v in ipairs(have) do
      if contains(winning, v) then
        wins = wins + 1
      end
    end

    out[#out + 1] = Card:new({
      id = string.gmatch(card, "[%d]+")(),
      wins = wins,
    })
  end
  return out
end

function p1(cards)
  local total = 0
  for _, card in ipairs(cards) do
    if card.wins > 0 then
      total = total + (2 ^ (card.wins - 1))
    end
  end
  return math.floor(total)
end

function main()
  local cards = load("assets/input.txt")
  print(p1(cards))
end

main()
