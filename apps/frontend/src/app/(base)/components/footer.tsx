import Link from "next/link";
import { GitHubLogoIcon } from "@radix-ui/react-icons";

export const Footer = () => {
  return (
    <footer
      className={`h-10 bg-blue-900 bg-opacity-20 relative before:content-[""] before:bg-blue-900 before:bg-opacity-20 before:h-full before:w-[10px] md:before:w-10 before:absolute before:top-0 before:-left-[10px] md:before:-left-10 after:content-[""] after:bg-blue-900 after:bg-opacity-20 after:h-full after:w-[10px] md:after:w-10 after:absolute after:top-0 after:-right-[10px] md:after:-right-10`}
    >
      <div className="size-full flex items-center justify-between">
        <Link
          href="https://github.com/urodstvo/book-shop"
          target="_blank"
          className="flex items-center gap-2"
        >
          <GitHubLogoIcon />
          github
        </Link>
        <span className="text-xs text-zinc-400">2024 © urodstvo</span>
      </div>
    </footer>
  );
};
