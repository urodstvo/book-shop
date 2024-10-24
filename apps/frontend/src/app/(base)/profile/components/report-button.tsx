"use client";

import { Button } from "@/components/ui/button";
import { API_URL } from "@/env";
import { FileScanIcon } from "lucide-react";

async function getReport(order_id: number) {
  const response = await fetch(`${API_URL}/orders/${order_id}/report`, {
    credentials: "include",
    mode: "cors",
  });
  const blob = await response.blob();
  const url = URL.createObjectURL(blob);
  window.open(url, "_blank");
}

export const ReportButton = ({ order_id }: { order_id: number }) => {
  return (
    <Button size="icon" variant="ghost" onClick={() => getReport(order_id)}>
      <FileScanIcon strokeWidth={2} />
    </Button>
  );
};
