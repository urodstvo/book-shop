"use client";
import * as React from "react";
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { API_URL } from "@/env";

const chartConfig = {
  rating: {
    label: "Средний рейтинг книг",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export default function BooksSection() {
  const [timeRange, setTimeRange] = React.useState<"30d" | "7d" | "90d">("30d");
  const [chartData, setChartData] = React.useState<{ date: string; rating: number }[]>([]);

  React.useEffect(() => {
    let daysToSubtract = 90;
    if (timeRange === "30d") {
      daysToSubtract = 30;
    } else if (timeRange === "7d") {
      daysToSubtract = 7;
    }
    const start = new Date();
    start.setDate(start.getDate() - daysToSubtract);
    const end = new Date();

    fetch(`${API_URL}/admin/books?start=${start.toISOString()}&end=${end.toISOString()}`, {
      method: "GET",
      credentials: "include",
    })
      .then((response) => response.json() as Promise<{ date: string; rating: number }[]>)
      .then((data) => {
        setChartData(data);
      });
  }, [timeRange]);

  return (
    <Card>
      <CardHeader className="flex items-center gap-2 space-y-0 border-b py-5 sm:flex-row">
        <div className="grid flex-1 gap-1 text-center sm:text-left">
          <CardTitle>Рейтинг книг </CardTitle>
          <CardDescription>График рейтинга книг за определенный период</CardDescription>
        </div>
        <Select
          value={timeRange}
          onValueChange={(value: "30d" | "7d" | "90d") => setTimeRange(value)}
        >
          <SelectTrigger className="w-[200px] rounded-lg sm:ml-auto" aria-label="Select a value">
            <SelectValue placeholder="Last 3 months" />
          </SelectTrigger>
          <SelectContent className="rounded-xl">
            <SelectItem value="90d" className="rounded-lg">
              Последние 3 месяца
            </SelectItem>
            <SelectItem value="30d" className="rounded-lg">
              Последние 30 дней
            </SelectItem>
            <SelectItem value="7d" className="rounded-lg">
              Посление 7 дней
            </SelectItem>
          </SelectContent>
        </Select>
      </CardHeader>
      <CardContent className="px-2 pt-4 sm:px-6 sm:pt-6">
        <ChartContainer config={chartConfig} className="aspect-auto h-[250px] w-full">
          <AreaChart data={chartData} accessibilityLayer>
            <defs>
              <linearGradient id="fillRating" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="var(--color-rating)" stopOpacity={0.8} />
                <stop offset="95%" stopColor="var(--color-rating)" stopOpacity={0.1} />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="date"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              minTickGap={32}
              tickFormatter={(value) => {
                const date = new Date(value);
                return date.toLocaleDateString("ru-RU", {
                  month: "short",
                  day: "numeric",
                });
              }}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              ticks={[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}
            />
            <ChartTooltip
              wrapperStyle={{ width: 200 }}
              cursor={false}
              content={
                <ChartTooltipContent
                  labelFormatter={(value) => {
                    return new Date(value).toLocaleDateString("ru-RU", {
                      month: "short",
                      day: "numeric",
                    });
                  }}
                  indicator="dot"
                />
              }
            />
            <Area
              dataKey="rating"
              type="bump"
              fill="url(#fillRating)"
              stroke="var(--color-rating)"
              stackId="a"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
