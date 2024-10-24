import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { InfoSection } from "./components/info";
import { OrdersSection } from "./components/orders";
import { RequestSection } from "./components/request-form";

export default function ProfilePage() {
  return (
    <Tabs defaultValue="info" className="gap-10 flex flex-col md:flex-row" orientation="vertical">
      <TabsList
        role="aside"
        className="w-full md:w-[300px] h-fit flex-col bg-transparent md:sticky md:top-[300px]"
      >
        <TabsTrigger
          className="w-full justify-start data-[state=active]:bg-primary data-[state=active]:text-white h-[32px] text-base px-10"
          value="info"
        >
          Информация
        </TabsTrigger>
        <TabsTrigger
          className="w-full justify-start data-[state=active]:bg-primary data-[state=active]:text-white text-base h-[32px] px-10"
          value="orders"
        >
          История заказов
        </TabsTrigger>
        <TabsTrigger
          className="w-full justify-start data-[state=active]:bg-primary data-[state=active]:text-white text-base h-[32px] px-10"
          value="request"
        >
          Попросить книгу
        </TabsTrigger>
      </TabsList>
      <TabsContent className="flex-1" value="info">
        <div className="bg-slate-300 w-full h-40 rounded-lg mb-10 md:block hidden" />
        <InfoSection />
      </TabsContent>
      <TabsContent className="flex-1" value="orders">
        <div className="bg-slate-300 w-full h-40 rounded-lg mb-10 md:block hidden" />
        <OrdersSection />
      </TabsContent>
      <TabsContent className="flex-1" value="request">
        <div className="bg-slate-300 w-full h-40 rounded-lg mb-10 md:block hidden" />
        <RequestSection />
      </TabsContent>
    </Tabs>
  );
}
