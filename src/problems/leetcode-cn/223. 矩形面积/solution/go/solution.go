func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    area, dx1, dx2, dy1, dy2 := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1), 0, 0, 0, 0
    if bx1 >= ax1 && bx1 <= ax2 {
        dx1 = bx1
    } else if ax1 >= bx1 && ax1 <= bx2 {
        dx1 = ax1
    }
    if bx2 >= ax1 && bx2 <= ax2 {
        dx2 = bx2
    } else if ax2 >= bx1 && ax2 <= bx2 {
        dx2 = ax2
    }
    if by1 >= ay1 && by1 <= ay2 {
        dy1 = by1
    } else if ay1 >= by1 && ay1 <= by2 {
        dy1 = ay1
    }
    if by2 >= ay1 && by2 <= ay2 {
        dy2 = by2
    } else if ay2 >= by1 && ay2 <= by2 {
        dy2 = ay2
    }
    return area-(dx2-dx1)*(dy2-dy1)
}