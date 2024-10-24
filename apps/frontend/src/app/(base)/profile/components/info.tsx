import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { PasswordInput } from "./password";
import { PaymentSection } from "./payment";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { API_URL } from "@/env";
import { User } from "models";
import { NameInput } from "./name";

export const InfoSection = async () => {
  if (!cookies().has("session_id")) redirect("/login");

  const response = await fetch(API_URL + "/users/me", {
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });
  if (!response.ok) {
    throw new Error("profile");
  }

  const user = (await response.json()) as User;

  return (
    <main className="flex gap-5 flex-col xl:flex-row xl:gap-40 items-center md:items-start">
      <section className="grid grid-cols-1 gap-10 py-10 w-[300px] place-content-center md:place-content-start">
        <div className="grid grid-col-1 w-full">
          <Label className="px-2">Имя пользователя</Label>
          <NameInput name={user.name} />
        </div>
        <div className="grid grid-col-1">
          <Label className="px-2">Логин</Label>
          <Input disabled defaultValue={user.login} placeholder="Иван" className="w-[300px]" />
        </div>
        <div className="grid grid-col-1 ">
          <Label className="px-2">Пароль</Label>
          <PasswordInput />
        </div>
        <div className="text-sm text-muted-foreground">
          Зарегистрирован с {new Date(user.created_at).toLocaleDateString("ru-RU")}
        </div>
      </section>
      <section className="flex-1">
        <h3 className="text-xl font-bold">Платежные методы</h3>
        <PaymentSection />
      </section>
    </main>
  );
};
