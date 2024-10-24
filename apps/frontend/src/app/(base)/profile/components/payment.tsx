import { MastercardIcon, MIRIcon, VisaIcon } from "./payment-method-icons";
import { API_URL } from "@/env";
import { Payment } from "models";
import { cookies } from "next/headers";
import { AddPaymentCard } from "./add-card";
import { DeletePaymentCard } from "./delete-card";

export const PaymentCard = ({
  card_number,
  cardholder_name,
  card_type,
  card_expired_at,
}: Payment) => {
  const date = new Date(card_expired_at);

  return (
    <div className="flex justify-center relative group">
      <div className="w-64 h-[120px] bg-gradient-to-r from-blue-700 via-blue-600 to-primary rounded-lg border py-2 px-5 shadow">
        <div className="flex justify-end items-center">
          {card_type === "visa" && <VisaIcon height={32} color="white" fill="white" />}
          {card_type === "mir" && <MIRIcon height={32} />}
          {card_type === "mastercard" && <MastercardIcon height={32} />}
        </div>
        <h1 className="w-full text-lg text-white">xxxx-xxxx-xxxx-{card_number.slice(-4)}</h1>
        <div className="flex flex-col justfiy-end text-white text-opacity-75 mt-2">
          <p className="font-bold text-xs">
            {date.getMonth() + 1} / {date.getFullYear()}
          </p>
          <h4 className="uppercase tracking-wider font-semibold text-xs">{cardholder_name}</h4>
        </div>
      </div>
      <DeletePaymentCard />
    </div>
  );
};

export const PaymentSection = async () => {
  const response = await fetch(API_URL + "/payments", {
    cache: "no-store",
    credentials: "include",
    headers: {
      Cookie: `session_id=${cookies().get("session_id")?.value}`,
    },
  });

  if (!response.ok) throw new Error("Failed to fetch payments");

  const payments = (await response.json()) as Payment[];

  return (
    <div className="w-full flex flex-wrap gap-5 mt-5 justify-center md:justify-start">
      <AddPaymentCard />
      {payments.map((payment) => (
        <PaymentCard key={payment.id} {...payment} />
      ))}
    </div>
  );
};
