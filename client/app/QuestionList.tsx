"use client";

import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import Link from "next/link";
import { use } from "react";

interface Question {
  id: number;
  no: number;
  name: string;
  subject: string;
  difficulty: string;
}

async function getQuestions() {
  const response = await fetch("http://localhost:8082/api/v1/questions");
  return response.json() as Promise<Question[]>;
}

const columnHelper = createColumnHelper<Question>();

const columns = [
  columnHelper.accessor("no", {
    header: () => "#",
    cell: (info) => info.renderValue(),
  }),
  columnHelper.accessor((row) => row.name, {
    id: "Title",
    cell: (info) => (
      <Link href={`/problem/${info.row.original.id}`}>{info.getValue()}</Link>
    ),
    header: () => <span>Last Name</span>,
    footer: (info) => info.column.id,
  }),
  columnHelper.accessor("difficulty", {
    header: () => "Difficulty",
    cell: (info) => info.renderValue(),
  }),
  columnHelper.accessor("subject", {
    header: () => "Subject",
    cell: (info) => info.renderValue(),
  }),
];

export default function QuestionList() {
  const questions = use(getQuestions());

  const table = useReactTable({
    data: questions,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div>
      <h1>Questions</h1>
      <table>
        <thead>
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <th key={header.id}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
