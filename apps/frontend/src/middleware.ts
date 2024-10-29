import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { API_URL } from "./env";
import { User } from "models";

export async function middleware(request: NextRequest) {
  const session = request.cookies.has("session_id");
  if (
    request.nextUrl.pathname.includes("/login") ||
    request.nextUrl.pathname.includes("/register")
  ) {
    if (session) return NextResponse.redirect(request.nextUrl.origin);
  }

  if (request.nextUrl.pathname.includes("/admin")) {
    if (!session) return NextResponse.error();
    else {
      const response = await fetch(API_URL + "/users/me", {
        cache: "no-store",
        headers: {
          Cookie: `session_id=${request.cookies.get("session_id")?.value}`,
        },
      });

      if (!response.ok) {
        return NextResponse.json({ message: "Internal server error" }, { status: 500 });
      }

      const user = (await response.json()) as User;

      if (user.role !== "admin")
        return NextResponse.json({ message: "Forbidden" }, { status: 403 });
    }
  }

  return NextResponse.next();
}
