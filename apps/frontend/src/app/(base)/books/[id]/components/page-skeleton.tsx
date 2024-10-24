import { AspectRatio } from "@/components/ui/aspect-ratio";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import { BadgeIcon } from "lucide-react";

export const PageSkeleton = () => {
  return (
    <main className="size-full">
      <section className="w-full grid md:grid-cols-[400px_800px] grid-rows-1 gap-5 md:gap-0 place-content-center">
        <div className="flex justify-center md:py-5">
          <div className="w-[300px] h-[400px] sticky top-10">
            <AspectRatio ratio={3 / 4}>
              <Skeleton className="size-full rounded-lg" />
              {/* <Image
            src="https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=160&dpr=2&q=80"
            alt="Placeholder image"
            fill
            className="h-full w-full rounded-lg object-cover"
            /> */}
              <div className="absolute top-2 right-2">
                <BadgeIcon
                  size={40}
                  fill="currentColor"
                  strokeWidth={1}
                  className="fill-yellow-700 stroke-yellow-700"
                />
              </div>
            </AspectRatio>
          </div>
        </div>
        <div className="flex flex-col gap-5 md:gap-10">
          <div className="flex flex-col gap-1">
            <h3>
              <Skeleton className=" h-[28px] w-[80%]" />
            </h3>
            <h5>
              <Skeleton className="w-1/2 h-[20px]" />
            </h5>
          </div>
          <Table className="w-full">
            <TableBody>
              <TableRow>
                <TableCell>Количество страниц</TableCell>
                <TableCell className="flex justify-end">
                  <Skeleton className="w-10 h-[20px]" />
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Издательство</TableCell>
                <TableCell className="flex justify-end">
                  <Skeleton className="w-40 h-[20px]" />
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Опубликовано</TableCell>
                <TableCell className="flex justify-end">
                  <Skeleton className="w-40 h-[20px]" />
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Количество заказов</TableCell>
                <TableCell className="flex justify-end">
                  <Skeleton className="w-10 h-[20px]" />
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Количество на складе</TableCell>
                <TableCell className="flex justify-end">
                  <Skeleton className="w-10 h-[20px]" />
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
          <span className="flex justify-end">
            <Skeleton className="w-[33%] h-[40px]" />
          </span>
          <div className="flex gap-5 ">
            <Button
              size="lg"
              className="rounded-full bg-indigo-700 hover:bg-indigo-800 text-white flex-[2] max-w-[600px]"
            >
              В корзину
            </Button>
            <Button size="lg" variant="outline" className="rounded-full max-w-[300px] flex-1">
              Демо
            </Button>
          </div>
          <div>
            <h6 className="text-lg mb-5">Аннотация</h6>
            <div className="mb-4 flex flex-col gap-1">
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-1/2 h-[24px]" />
            </div>
            <div className="mb-4 flex flex-col gap-1">
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-1/3 h-[24px]" />
            </div>
            <div className="mb-4 flex flex-col gap-1">
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-full h-[24px]" />
              <Skeleton className="w-[80%] h-[24px]" />
            </div>
          </div>
        </div>
      </section>
    </main>
  );
};
