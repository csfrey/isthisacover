"use client";
import { useQuery } from "@tanstack/react-query";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Image from "next/image";
import axios from "axios";
import { API_BASE } from "@/app/constants";

export default function Home() {
  const test = useQuery({
    queryKey: ["Test"],
    queryFn: async () => {
      const response = await axios.get(`${API_BASE}/`);
      return response.data;
    },
  });

  return (
    <main className="min-h-screen flex items-center justify-center dark:text-white dark:bg-gray-950">
      <div className="mb-12 flex flex-col gap-4 w-[500px]">
        <div className="text-4xl w-full text-center">Is This a Cover?</div>
        <div className="flex gap-2">
          <Input placeholder="Search for a song on Spotify..." autoFocus />
          <Button className="bg-green-500 hover:bg-green-600 active:bg-green-700">
            GO
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className="size-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="m8.25 4.5 7.5 7.5-7.5 7.5"
              />
            </svg>
          </Button>
        </div>

        <div>API says:</div>
        <div>{test.isFetching ? "Loading..." : test.data?.message}</div>
      </div>
    </main>
  );
}
