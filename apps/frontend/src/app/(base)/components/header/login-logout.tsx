import { Button } from "@/components/ui/button";
import { LogoutButton } from "./logout";
import { cookies } from "next/headers";
import Link from "next/link";

export const LoginLogoutButton = () => {
  const cookieStore = cookies();
  const isAuth = cookieStore.get("session_id");

  if (isAuth) {
    return <LogoutButton />;
  }

  return (
    <Button variant="default" className="w-40 rounded-full text-white" asChild>
      <Link href="/login">Войти</Link>
    </Button>
  );
};
