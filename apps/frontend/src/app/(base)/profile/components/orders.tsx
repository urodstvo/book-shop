import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { API_URL } from "@/env";
import { Order, Payment } from "models";
import { cookies } from "next/headers";
import { ReportButton } from "./report-button";
import { CancelButton } from "./cancel-button";

export const OrdersSection = async () => {
  const response = await fetch(API_URL + "/orders", {
    cache: "no-store",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });

  if (!response.ok) throw new Error("Failed to fetch invoices.");

  const { orders } = (await response.json()) as { orders: (Order & { payment: Payment })[] };

  const formattedCardNumber = (cardNumber: string) => {
    return cardNumber
      .replaceAll("-", "")
      .replace(/(\d{4})/g, "$1-")
      .slice(0, -1);
  };
  return (
    <main>
      <h3 className="text-2xl font-bold">История заказов</h3>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>№</TableHead>
            <TableHead>Статус</TableHead>
            <TableHead className="w-[200px]">Карта</TableHead>
            <TableHead>Дата</TableHead>
            <TableHead className="text-right">Сумма</TableHead>
            <TableHead className="text-center">Отчет</TableHead>
            <TableHead className="text-center">Отмена</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {orders?.map((order) => (
            <TableRow key={order.id}>
              <TableCell className="font-medium text-sm">{order.id}</TableCell>
              <TableCell className="text-sm">
                {order.status === "pending" && "Ожидается"}
                {order.status === "delivered" && "Доставлен"}
                {order.status === "cancelled" && "Отменен"}
              </TableCell>
              <TableCell className="text-sm w-[200px] whitespace-nowrap">
                {formattedCardNumber(order.payment.card_number)}
              </TableCell>
              <TableCell className="text-sm">{new Date().toLocaleDateString()}</TableCell>
              <TableCell className="text-right text-smt">{order.price} ₽</TableCell>
              <TableCell className="text-center">
                <ReportButton order_id={order.id} />
              </TableCell>
              <TableCell className="text-center">
                {order.status === "pending" && <CancelButton order_id={order.id} />}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </main>
  );
};
