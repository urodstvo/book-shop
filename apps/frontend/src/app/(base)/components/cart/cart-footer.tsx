"use client";

import { Label } from "@/components/ui/label";

import { PaymentSelect } from "./payment-select";
import { Cart, Payment } from "models";
import { Button } from "@/components/ui/button";
import { SheetFooter } from "@/components/ui/sheet";
import { startTransition, useEffect, useState } from "react";
import { API_URL } from "@/env";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const createOrder = async (paymentId: number) => {
  const carts = await getCarts();

  const response = await fetch(`${API_URL}/orders`, {
    method: "POST",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      payment_id: paymentId,
      books: carts.map((cart) => ({ book_id: cart.item.book_id, amount: cart.item.quantity })),
    }),
  });
  if (!response.ok) {
    toast.error("Не удалось создать заказ");
    throw new Error("Failed to create order");
  }

  toast.success("Заказ создан");
};

async function getCarts() {
  const response = await fetch(API_URL + "/carts", {
    cache: "no-store",
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch carts");
  }

  return (await response.json()) as { item: Cart }[];
}
export const CartFooter = ({ payments }: { payments: Payment[] }) => {
  const router = useRouter();
  const [selectedPayment, setSelectedPayment] = useState<number | null>(null);
  const [isCartEmpty, setIsCartEmpty] = useState(false);

  useEffect(() => {
    getCarts().then((carts) => {
      setIsCartEmpty(carts.length === 0);
    });
  }, []);

  return (
    <SheetFooter>
      <div className="flex flex-col gap-5 w-full">
        <div className="flex justify-between items-center">
          <Label>Способ оплаты</Label>
          <PaymentSelect payments={payments} setPayment={setSelectedPayment} />
        </div>
        <Button
          className="text-white w-full"
          disabled={!selectedPayment || isCartEmpty}
          onClick={() => {
            if (selectedPayment) {
              createOrder(selectedPayment);

              startTransition(() => router.refresh());
            }
          }}
        >
          Оформить заказ
        </Button>
      </div>
    </SheetFooter>
  );
};
