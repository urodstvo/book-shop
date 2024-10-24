import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { DashboardIcon } from "@radix-ui/react-icons";
import { User } from "models";
import { cookies } from "next/headers";
import Link from "next/link";

export const Dashboard = async () => {
  if (!cookies().has("session_id")) return null;

  const session = cookies().get("session_id");

  const response = await fetch(API_URL + "/users/me", {
    credentials: "include",
    headers: {
      Cookie: `session_id=${session?.value}`,
    },
  });

  if (!response.ok) return null;

  const user = (await response.json()) as User;

  if (user.role !== "admin") return null;

  return (
    <Button asChild size="icon" variant="ghost" className="size-10 [&_svg]:size-[20px]">
      <Link href="/admin/dashboard">
        <DashboardIcon strokeWidth={2} />
      </Link>
    </Button>
  );
};
