"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { startTransition, useState } from "react";
import { toast } from "sonner";
import { API_URL } from "@/env";
import { useRouter } from "next/navigation";

const formSchema = z
  .object({
    firstName: z
      .string({
        required_error: "Поле обязательно",
      })
      .min(2, { message: "Поле обязательно" })
      .max(50, { message: "Поле должно содержать не более 50 символов" }),
    lastName: z
      .string({
        required_error: "Поле обязательно",
      })
      .min(2, { message: "Поле обязательно" })
      .max(50, { message: "Поле должно содержать не более 50 символов" }),
    login: z
      .string({
        required_error: "Поле обязательно",
      })
      .min(2, { message: "Поле обязательно" })
      .max(50, { message: "Поле должно содержать не более 50 символов" }),
    password: z
      .string({
        required_error: "Поле обязательно",
      })
      .min(6, { message: "Поле должно содержать не менее 6 символов" })
      .max(50, { message: "Поле должно содержать не более 50 символов" }),
    repeatPassword: z
      .string({
        required_error: "Поле обязательно",
      })
      .max(50, { message: "Поле должно содержать не более 50 символов" }),
  })
  .refine((data) => data.password === data.repeatPassword, {
    message: "Пароли не совпадают",
    path: ["repeatPassword"],
  });

export default function RegisterPage() {
  const router = useRouter();
  const [isPending, setIsPending] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsPending(true);
    const data = JSON.stringify({
      name: values.lastName + " " + values.firstName,
      login: values.login,
      password: values.password,
    });

    const response = await fetch(API_URL + "/auth/register", {
      method: "POST",
      mode: "cors",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: data,
    });

    setIsPending(false);
    if (!response.ok) {
      toast.error("Ошибка при регистрации");
      return;
    }

    toast.success("Аккаунт создан");
    startTransition(() => {
      router.push("/");
      router.refresh();
    });
  }

  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-xl">Регистрация</CardTitle>
        <CardDescription>Введите данные для создания аккаунта</CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <div className="grid gap-4">
              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="firstName"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Имя</FormLabel>
                      <FormControl>
                        <Input placeholder="Иван" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="lastName"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Фамилия</FormLabel>
                      <FormControl>
                        <Input placeholder="Иванов" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <FormField
                control={form.control}
                name="login"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Логин</FormLabel>
                    <FormControl>
                      <Input placeholder="abobameister" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Пароль</FormLabel>
                    <FormControl>
                      <Input placeholder="" {...field} type="password" />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="repeatPassword"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Повторите пароль</FormLabel>
                    <FormControl>
                      <Input placeholder="" {...field} type="password" />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <Button type="submit" className="w-full text-white" disabled={isPending}>
                {isPending ? "Создание..." : "Создать аккаунт"}
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Уже есть аккаунт?{" "}
              <Link href="/login" className="underline">
                Авторизуйтесь
              </Link>
            </div>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
