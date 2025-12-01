dx = 94.0
dx_ = 22.0
x = 8400.0

dy = 34.0
dy_ = 67.0
y = 5400.0

a = ((x/dx) - ((dx_*y)/(dy*dy_))) * (dy_/(dy_-dy))

b = (y/dy_)-((dy/dy_)*a)

print(a, b)
