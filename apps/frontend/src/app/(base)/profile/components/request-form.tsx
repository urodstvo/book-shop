"use client";

import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
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
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { API_URL } from "@/env";
import { toast } from "sonner";
import { useState } from "react";

const schema = z.object({
  author: z
    .string({
      required_error: "Поле обязательно",
    })
    .min(1, { message: "Поле обязательно" }),
  name: z
    .string({
      required_error: "Поле обязательно",
    })
    .min(1, { message: "Поле обязательно" }),
  comment: z.string().optional(),
});

const saveRequest = async (book_name: string, comment?: string) => {
  const data: { [key: string]: string } = { book_name };
  if (comment) data.comment = comment;

  const response = await fetch(API_URL + "/books/request", {
    method: "POST",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    toast.error("Не удалось отправить запрос");
    throw new Error("Failed to save request");
  }

  toast.success("Запрос отправлен");
};

export const RequestSection = () => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      author: "",
      name: "",
      comment: "",
    },
  });

  return (
    <main className="flex justify-center">
      <Card className="w-full lg:w-[600px]">
        <CardHeader>
          <CardTitle>Форма запроса книги, не имеющейся в каталоге</CardTitle>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit((v) => {
                setIsSubmitting(true);
                saveRequest(`${v.author.trim()}: ${v.name.trim()}`, v.comment?.trim())
                  .then(() =>
                    form.reset({
                      author: "",
                      name: "",
                      comment: "",
                    })
                  )
                  .finally(() => setIsSubmitting(false));
              })}
              className="flex flex-col gap-5"
            >
              <FormField
                control={form.control}
                name="author"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Имя автора</FormLabel>
                    <FormControl>
                      <Input placeholder="Имя автора" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Название книги</FormLabel>
                    <FormControl>
                      <Input placeholder="Название книги" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="comment"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Комментарий к запросу</FormLabel>
                    <FormControl>
                      <Textarea placeholder="Введите комментарий" {...field} />
                    </FormControl>
                    <FormDescription>Это поле не обязательно для заполнения</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <CardFooter>
                <Button className="w-full" type="submit" disabled={isSubmitting}>
                  {isSubmitting ? "Отправка..." : "Отправить запрос"}
                </Button>
              </CardFooter>
            </form>
          </Form>
        </CardContent>
      </Card>
    </main>
  );
};
