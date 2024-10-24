import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Payment } from "models";
import { Dispatch, SetStateAction } from "react";

export const PaymentSelect = ({
  payments,
  setPayment,
}: {
  payments: Payment[];
  setPayment: Dispatch<SetStateAction<number | null>>;
}) => {
  return (
    <Select onValueChange={(v) => setPayment(Number(v))}>
      <SelectTrigger className="max-w-[200px]">
        <SelectValue placeholder="Выберите способ" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {payments.map((payment) => (
            <SelectItem key={payment.id} value={payment.id.toString()}>
              xxxx-xxx-xxxx-{payment.card_number.slice(-4)}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
