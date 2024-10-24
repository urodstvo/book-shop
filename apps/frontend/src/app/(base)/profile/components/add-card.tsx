"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTrigger,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { PlusIcon } from "lucide-react";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

import * as validator from "card-validator";
import { startTransition, useEffect, useState } from "react";
import { PatternFormat } from "react-number-format";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { API_URL } from "@/env";

const schema = z.object({
  cardholder: z.string({
    required_error: "Поле обязательно",
  }),
  card_number: z
    .string({
      required_error: "Поле обязательно",
    })
    .refine((v) => validator.number(v.replaceAll("-", "").replaceAll("*", "")).isValid, {
      message: "Неверная карта",
    }),
  expired_at: z
    .string({
      required_error: "Поле обязательно",
    })
    .refine((v) => validator.expirationDate(v).isValid, {
      message: "Неверная дата",
    }),
});

const TypeConverter = (type: "mastercard" | "visa" | "mir" | "") => {
  switch (type) {
    case "mastercard":
      return "MasterCard";
    case "visa":
      return "Visa";
    case "mir":
      return "МИР";
    default:
      return "";
  }
};

const AddPaymentCardForm = () => {
  const [type, setType] = useState("");
  const [isPending, setIsPending] = useState(false);
  const router = useRouter();

  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    mode: "onBlur",
    defaultValues: {
      cardholder: "",
      card_number: "",
      expired_at: "",
    },
  });

  const cardNumber = form.watch("card_number");

  useEffect(() => {
    const num = cardNumber.replaceAll("-", "").replaceAll("*", "");
    setType(validator.number(num).card?.type || "");
  }, [cardNumber]);

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit((v) => {
          const date = v.expired_at.split("/").map((el) => el.trim());
          if (date[1].length === 2) date[1] = `20${date[1]}`;

          const dateStr = `${date[1]}-${date[0]}-01T00:00:00.000Z`;

          const data = JSON.stringify({
            cardholder: v.cardholder,
            card_number: v.card_number.replaceAll("-", "").replaceAll("*", ""),
            expired_at: new Date(dateStr).toISOString(),
            type,
          });

          setIsPending(true);
          fetch(API_URL + "/payments", {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: data,
          })
            .then((res) => {
              setIsPending(false);
              if (res.ok) {
                toast.success("Карта добавлена");
                startTransition(() => {
                  router.refresh();
                  form.reset();
                });
              }
            })
            .catch(() => {
              setIsPending(false);
              toast.error("Ошибка при добавлении платежной карты");
            });
        })}
        className="flex flex-col gap-4"
      >
        <FormField
          control={form.control}
          name="card_number"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Номер карты</FormLabel>
              <FormControl>
                <PatternFormat
                  format="####-####-####-####"
                  {...field}
                  onValueChange={({ value }) => field.onChange(value)}
                  customInput={Input}
                  mask={"*"}
                  allowEmptyFormatting
                />
              </FormControl>
              <FormDescription>Номер карты без пробелов</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Input value={TypeConverter(type as "mastercard" | "visa" | "mir" | "")} disabled />
        <FormField
          control={form.control}
          name="cardholder"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Имя на карте</FormLabel>
              <FormControl>
                <Input placeholder="Ivanov Ivan" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="expired_at"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Действует до</FormLabel>
              <FormControl>
                <Input placeholder="09/2024" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <DialogFooter>
          <Button type="submit" disabled={isPending}>
            Добавить
          </Button>
        </DialogFooter>
      </form>
    </Form>
  );
};

export const AddPaymentCard = () => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          size={null}
          variant="ghost"
          className="[&_svg]:size-[48px] w-64 h-[120px] border rounded-lg"
        >
          <PlusIcon size={64} color="gray" />
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Добавление платежной карты</DialogTitle>
          <DialogDescription>Заполните форму для добавления платежной карты</DialogDescription>
        </DialogHeader>
        <AddPaymentCardForm />
      </DialogContent>
    </Dialog>
  );
};
