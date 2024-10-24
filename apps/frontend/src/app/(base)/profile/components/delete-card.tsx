"use client";

import { Button } from "@/components/ui/button";
import { TrashIcon } from "lucide-react";

export const DeletePaymentCard = () => {
  return (
    <Button
      size="icon"
      variant="secondary"
      className="absolute -top-2 right-[-16px] [&_svg]:size-4 size-[32px] group-hover:opacity-100 opacity-0 transition-opacity"
    >
      <TrashIcon color="black" />
    </Button>
  );
};
