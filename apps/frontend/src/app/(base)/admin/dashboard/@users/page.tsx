"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useState } from "react";
import React from "react";
import { API_URL } from "@/env";
import { Order } from "models";

type response = {
  orders: (Order & {
    user_name: string;
    user_login: string;
    book_count: number;
  })[];
  unique_users: number;
  total_price: number;
  total_ordered_books: number;
};

export default function UsersSection() {
  const [timeRange, setTimeRange] = useState<"30d" | "7d" | "90d">("30d");
  const [data, setData] = React.useState<response>({
    orders: [],
    unique_users: 0,
    total_price: 0,
    total_ordered_books: 0,
  });

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

    fetch(`${API_URL}/admin/users?start=${start.toISOString()}&end=${end.toISOString()}`, {
      method: "GET",
      credentials: "include",
    })
      .then((response) => response.json() as Promise<response>)
      .then((data) => {
        setData(data);
      });
  }, [timeRange]);

  return (
    <section>
      <h3 className="font-bold text-lg">Заказы</h3>
      <p className="text-muted-foreground text-sm">
        Количество оформленных заказов пользователями и итоговая сумма заказов за определенный
        период
      </p>
      <div className="mt-5 mb-5">
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
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>№ заказа</TableHead>
            <TableHead>Пользователь</TableHead>
            <TableHead>Книг в заказе</TableHead>
            <TableHead>Сумма</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data.orders.map((el) => (
            <TableRow key={el.id}>
              <TableCell className="font-medium text-sm">{el.id}</TableCell>
              <TableCell className="font-medium text-sm">{el.user_name || el.user_login}</TableCell>
              <TableCell className="font-medium text-sm">{el.book_count}</TableCell>
              <TableCell className="font-medium text-sm">{el.price} ₽</TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell>Всего</TableCell>
            <TableCell>{data.unique_users}</TableCell>
            <TableCell>{data.total_ordered_books}</TableCell>
            <TableCell>{data.total_price} ₽</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    </section>
  );
}
