import { Button } from "@/components/ui/button";
import { UserIcon } from "lucide-react";
import { cookies } from "next/headers";
import Link from "next/link";

export const Profile = () => {
  if (!cookies().has("session_id")) return null;

  return (
    <Button asChild size="icon" variant="ghost" className="size-10 [&_svg]:size-[20px]">
      <Link href="/profile">
        <UserIcon strokeWidth={2} />
      </Link>
    </Button>
  );
};
