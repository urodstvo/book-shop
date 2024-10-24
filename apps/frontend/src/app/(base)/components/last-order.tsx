import { API_URL } from "@/env";
import { Order } from "models";
import { cookies } from "next/headers";

export const OrderSection = async () => {
  if (!cookies().has("session_id")) return null;
  const response = await fetch(API_URL + "/orders", {
    cache: "no-store",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });

  if (!response.ok) return null;

  const { orders } = (await response.json()) as { orders: Order[] };

  if (orders.length === 0) return null;

  const lastOrder = orders.sort(
    (o1, o2) => new Date(o2.updated_at).getTime() - new Date(o1.updated_at).getTime()
  )[0];

  return (
    <section className="flex justify-center mb-20">
      <div className="bg-accent px-[30px] py-4 rounded-xl w-full max-w-[1080px] flex justify-between item-center">
        <h3 className="text-xl">
          Последний заказ: №{lastOrder.id} от{" "}
          {new Intl.DateTimeFormat("ru").format(new Date(lastOrder.created_at))}
        </h3>
        <p>
          Статус: {lastOrder.status === "pending" && "Ожидается"}
          {lastOrder.status === "delivered" && "Доставлен"}
          {lastOrder.status === "cancelled" && "Отменен"}
        </p>
      </div>
    </section>
  );
};
