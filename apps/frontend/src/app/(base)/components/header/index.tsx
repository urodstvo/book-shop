import { ShoppingCart } from "../cart";
import { Search } from "./search";
import { HomeLink } from "./home";
import { Dashboard } from "./dashboard";
import { LoginLogoutButton } from "./login-logout";
import { Profile } from "./profile";

export const Header = () => {
  return (
    <header className="grid grid-cols-1 grid-rows-2 gap-2 md:grid-rows-1 md:grid-cols-3 md:h-20 py-5 w-full grid-flow-col-dense">
      <div className="lcol flex w-fit row-[1] md:row-auto md:w-full">
        <HomeLink />
      </div>
      <div className="mcol flex items-center row-[2] col-span-2 md:col-span-1 md:row-auto">
        <Search />
      </div>
      <nav className="rcol size-full flex items-center justify-end gap-1 md:gap-5 row-[1] md:row-auto w-fit md:w-full">
        <Dashboard />
        <Profile />
        <ShoppingCart />
        <LoginLogoutButton />
      </nav>
    </header>
  );
};
