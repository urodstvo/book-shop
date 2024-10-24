"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { useRouter } from "next/navigation";
import { startTransition } from "react";

export const LogoutButton = () => {
  const router = useRouter();
  return (
    <Button
      variant="default"
      className="w-40 rounded-full text-white"
      onClick={() => {
        fetch(API_URL + "/auth/logout", {
          method: "POST",
          credentials: "include",
        }).then(() => {
          startTransition(() => {
            router.refresh();
          });
        });
      }}
    >
      Выйти
    </Button>
  );
};
