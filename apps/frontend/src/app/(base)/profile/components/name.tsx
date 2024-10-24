"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { API_URL } from "@/env";
import { PencilIcon, SaveIcon, XIcon } from "lucide-react";
import { useState } from "react";
import { toast } from "sonner";

async function changeName(name: string) {
  const response = await fetch(API_URL + "/users/me", {
    method: "PUT",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name }),
  });

  if (!response.ok) {
    if (response.status === 400) {
      toast.error("Ошибка при изменении имени");
    }
    throw new Error("Failed to change name");
  }
}

export const NameInput = ({ name }: { name: string }) => {
  const [value, setValue] = useState(name);
  const [disabled, setDisabled] = useState(true);

  return (
    <div className="flex gap-1 items-center w-[300px]">
      <Input
        disabled={disabled}
        value={value}
        placeholder="Иван"
        className="max-w-[300px] flex-1"
        onChange={(event) => setValue(event.target.value)}
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
              changeName(value).then(() => {
                setDisabled(!disabled);
                toast.success("Имя успешно изменено");
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
