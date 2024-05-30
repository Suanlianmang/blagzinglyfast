package pages

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Count struct {
    Count int
}

func Index(c echo.Context) error {
    count := c.Get("count").(*Count)
    return c.Render(http.StatusOK, "index", count)
}

func Increment(c echo.Context) error {
    count := c.Get("count").(*Count)
    count.Count += 1;
    return c.Render(http.StatusOK, "count", count)
}
