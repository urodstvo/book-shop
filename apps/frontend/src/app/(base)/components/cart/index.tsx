import { Button } from "@/components/ui/button";

import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { Table, TableBody, TableCell, TableFooter, TableRow } from "@/components/ui/table";
import { API_URL } from "@/env";
import { ShoppingCartIcon } from "lucide-react";
import { Cart, Payment } from "models";
import { cookies } from "next/headers";
import { DeleteButton } from "./remove-book";
import { PlusButton } from "./add-button";
import { MinusButton } from "./remove-button";
import { QuantityInput } from "./cart-input";
import { CartFooter } from "./cart-footer";

export const ShoppingCart = async () => {
  if (!cookies().has("session_id")) {
    return null;
  }

  const cartsResponse = await fetch(API_URL + "/carts", {
    cache: "no-store",
    mode: "cors",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });
  const carts = (await cartsResponse.json()) as {
    item: Cart;
    stock_count: number;
    book_name: string;
    book_author: string;
    book_price: number;
  }[];

  const PaymentsResponse = await fetch(API_URL + "/payments", {
    credentials: "include",
    cache: "no-store",
    mode: "cors",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });
  const payments = (await PaymentsResponse.json()) as Payment[];

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button size="icon" variant="ghost" className="size-10 [&_svg]:size-[20px]">
          <ShoppingCartIcon strokeWidth={2} />
        </Button>
      </SheetTrigger>
      <SheetContent className="w-full sm:max-w-none md:w-[500px]">
        <SheetHeader>
          <SheetTitle>Корзина</SheetTitle>
        </SheetHeader>

        <Table className="mt-5 mb-5">
          <TableBody>
            {carts.map((cart) => (
              <TableRow key={cart.item.book_id} className="">
                <TableCell className="w-full">{cart.book_name}</TableCell>
                <TableCell className="whitespace-nowrap">{cart.book_price} ₽</TableCell>
                <TableCell className="gap-1 flex justify-end w-full">
                  <PlusButton
                    book_id={cart.item.book_id}
                    session={cookies().get("session_id")?.value}
                    quantity={cart.item.quantity}
                  />
                  <QuantityInput
                    quantity={cart.item.quantity}
                    session={cookies().get("session_id")?.value}
                    stock_count={cart.stock_count}
                    book_id={cart.item.book_id}
                  />
                  <MinusButton
                    book_id={cart.item.book_id}
                    session={cookies().get("session_id")?.value}
                    quantity={cart.item.quantity}
                  />
                  <DeleteButton
                    session={cookies().get("session_id")?.value}
                    book_id={cart.item.book_id}
                  />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
          <TableFooter>
            <TableRow>
              <TableCell colSpan={2}>Итого</TableCell>
              <TableCell className="text-end">
                {carts.reduce((a, b) => a + b.book_price * b.item.quantity, 0)} ₽
              </TableCell>
            </TableRow>
          </TableFooter>
        </Table>
        <CartFooter payments={payments} />
      </SheetContent>
    </Sheet>
  );
};
