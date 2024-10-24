"use client";

import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { CheckIcon, ListFilterIcon } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { useCallback, useEffect, useState } from "react";
import { cn } from "@/lib/utils";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Genre } from "models";
import { API_URL } from "@/env";

const GenrePicker = () => {
  const [genres, setGenres] = useState<Genre[]>([]);
  const [selectedGenres, setSelectedGenres] = useState<Genre[]>([]);

  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  useEffect(() => {
    fetch(API_URL + "/genres", {
      mode: "cors",
    })
      .then((response) => response.json() as Promise<Genre[]>)
      .then((data) => {
        setGenres(data);
      })
      .catch(() => {});
  }, []);

  useEffect(() => {
    const params = new URLSearchParams(searchParams);
    setSelectedGenres(
      genres.filter((genre) => params.getAll("genre").includes(genre.id.toString()))
    );
  }, [genres, searchParams]);

  useEffect(() => {
    const params = new URLSearchParams(searchParams);
    params.getAll("genre").forEach(() => params.delete("genre"));

    selectedGenres.forEach((genre) => params.append("genre", genre.id.toString()));

    replace(`${pathname}?${params.toString()}`, {
      scroll: false,
    });
  }, [pathname, replace, searchParams, selectedGenres]);

  const toggleGenre = useCallback(
    (genre: { id: number; name: string }) => {
      if (selectedGenres.some((g) => g.id === genre.id)) {
        setSelectedGenres((prev) => prev.filter((g) => g.id !== genre.id));
      } else {
        setSelectedGenres((prev) => [...prev, genre]);
      }
    },
    [selectedGenres]
  );
  return (
    <Command>
      <CommandInput placeholder="Поиск жанров..." />
      <CommandList>
        <CommandEmpty>Нет найденных жанров</CommandEmpty>
        <CommandGroup>
          <ScrollArea className="h-[200px] w-full pr-4">
            {genres.map((genre) => (
              <CommandItem key={genre.id} value={genre.name} onSelect={() => toggleGenre(genre)}>
                <CheckIcon
                  className={cn(
                    "mr-2 h-4 w-4",
                    selectedGenres.some((g) => g.id === genre.id) ? "opacity-100" : "opacity-0"
                  )}
                />
                {genre.name}
              </CommandItem>
            ))}
          </ScrollArea>
        </CommandGroup>
      </CommandList>
      <span className="text-sm text-muted-foreground">Выбрано жанров: {selectedGenres.length}</span>
    </Command>
  );
};

const TypeToggler = () => {
  const [searchType, setSearchType] = useState<"name" | "author">("name");

  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  useEffect(() => {
    const params = new URLSearchParams(searchParams);
    const type = params.get("type") as "name" | "author" | undefined;
    if (type) setSearchType(type);
  }, [searchParams]);

  useEffect(() => {
    const params = new URLSearchParams(searchParams);
    params.set("type", searchType);

    replace(`${pathname}?${params.toString()}`, {
      scroll: false,
    });
  }, [pathname, replace, searchType, searchParams]);

  return (
    <div>
      <p className="text-sm px-2">Текстовый поиск</p>
      <div className="flex gap-1 bg-slate-300 bg-opacity-50 p-1 rounded-full items-center">
        <Button
          size="sm"
          variant={searchType === "name" ? "default" : "outline"}
          className={cn(
            "flex-1 px-2 rounded-full text-sm flex items-center text-center hover:bg-zinc-700 bg-transparent bg-opacity-80 border-transparent border",
            searchType === "name" && "bg-primary hover:bg-primary/90"
          )}
          onClick={() => setSearchType("name")}
        >
          По названию
        </Button>
        <Button
          size="sm"
          variant={searchType === "author" ? "default" : "outline"}
          className={cn(
            "flex-1 px-2 rounded-full text-sm flex items-center text-center hover:bg-zinc-700 bg-transparent bg-opacity-80 border-transparent border",
            searchType === "author" && "bg-primary hover:bg-primary/90"
          )}
          onClick={() => setSearchType("author")}
        >
          По автору
        </Button>
      </div>
    </div>
  );
};

export const Search = () => {
  const pathname = usePathname();
  const isHomePage = pathname === "/";

  const searchParams = useSearchParams();
  const { replace } = useRouter();

  return (
    isHomePage && (
      <div className="size-full flex items-center h-10 bg-primary p-[6px] rounded-full gap-2">
        <Input
          autoFocus
          name="search-book-store"
          type="text"
          placeholder="..."
          className="w-full px-4 border border-primary/10 rounded-full text-base flex-1 h-[28px] text-black bg-secondary"
          onChange={(event) => {
            const params = new URLSearchParams(searchParams);
            params.set("q", event.target.value);
            replace(`${pathname}?${params.toString()}`, {
              scroll: false,
            });
          }}
        />
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="secondary"
              size="icon"
              className="rounded-full size-[28px] dark:bg-zinc-200 dark:hover:bg-zinc-300 "
            >
              <ListFilterIcon strokeWidth={2} />
            </Button>
          </PopoverTrigger>
          <PopoverContent align="end" sideOffset={16} className="min-w-[300px] md:w-[300px]">
            <div className="flex flex-col gap-2 ">
              <b className="mb-2">Фильтр</b>
              <TypeToggler />
              <GenrePicker />
            </div>
          </PopoverContent>
        </Popover>
      </div>
    )
  );
};
