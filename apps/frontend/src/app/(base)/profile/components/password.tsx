"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { API_URL } from "@/env";
import { PencilIcon, SaveIcon, XIcon } from "lucide-react";
import { useState } from "react";
import { toast } from "sonner";

async function changePassword(password: string) {
  const response = await fetch(API_URL + "/users/me", {
    method: "PUT",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ password }),
  });

  if (!response.ok) {
    if (response.status === 400) {
      toast.error("Ошибка при изменении пароля");
    }
    throw new Error("Failed to change password");
  }
}

export const PasswordInput = () => {
  const [disabled, setDisabled] = useState(true);
  const [value, setValue] = useState("your-password");

  return (
    <div className="flex gap-1 items-center w-[300px]">
      <Input
        disabled={disabled}
        value={value}
        onChange={(event) => setValue(event.target.value)}
        type={disabled ? "password" : "text"}
        className="max-w-[300px] flex-1"
      />
      {disabled ? (
        <Button size="icon" variant="ghost" onClick={() => setDisabled(!disabled)}>
          <PencilIcon strokeWidth={2} />
        </Button>
      ) : (
        <>
          <Button
            size="icon"
            variant="ghost"
            onClick={() =>
              changePassword(value).then(() => {
                setDisabled(!disabled);
                toast.success("Пароль успешно изменен");
              })
            }
          >
            <SaveIcon strokeWidth={2} />
          </Button>
          <Button
            size="icon"
            variant="ghost"
            onClick={() => {
              setDisabled(!disabled);
              setValue("your-password");
            }}
          >
            <XIcon strokeWidth={2} />
          </Button>
        </>
      )}
    </div>
  );
};
