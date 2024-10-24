import { Button } from "@/components/ui/button";
import { BookUserIcon } from "lucide-react";
import Link from "next/link";

export const HomeLink = () => {
  return (
    <Button asChild variant={null} size="icon" className="[&_svg]:size-[24px] size-10">
      <Link href="/" className="text-base font-medium">
        <BookUserIcon strokeWidth={2} />
      </Link>
    </Button>
  );
};
