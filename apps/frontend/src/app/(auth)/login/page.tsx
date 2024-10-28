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
import { API_URL } from "@/env";
import { toast } from "sonner";
import { startTransition, useState } from "react";
import { useRouter } from "next/navigation";

const formSchema = z.object({
  login: z.string({
    required_error: "Поле обязательно",
  }),
  password: z.string({
    required_error: "Поле обязательно",
  }),
});

export default function LoginPage() {
  const [isPending, setIsPending] = useState(false);
  const router = useRouter();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const data = JSON.stringify(values);

    setIsPending(true);

    const response = await fetch(API_URL + "/auth/login", {
      method: "POST",
      mode: "cors",
      credentials: "include",
      body: data,
    });

    setIsPending(false);
    if (!response.ok) {
      toast.error("Неверные логин или пароль");
      return;
    }

    toast.success("Авторизация прошла успешно");
    startTransition(() => {
      router.push("/");
      router.refresh();
    });
  }

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Авторизация</CardTitle>
        <CardDescription>Введите логин и пароль для входа</CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <div className="grid gap-4">
              <FormField
                control={form.control}
                name="login"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Логин</FormLabel>
                    <FormControl>
                      <Input placeholder="aboba" {...field} />
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
                      <Input
                        placeholder=""
                        type="password"
                        {...field}
                        autoComplete="current-password"
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <Button type="submit" className="w-full text-white" disabled={isPending}>
                {isPending ? "Авторизация..." : "Войти"}
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Еще нет аккаунта?{" "}
              <Link href="/register" className="underline">
                Зарегистрироваться
              </Link>
            </div>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
