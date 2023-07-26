import { z } from 'zod';

export const ZSandbox = z.object({
  id: z.number(),
  name: z.string(),
  language: z.string(),
  code: z.string(),
  status: z.string(),
  stdout: z.string(),
  stderr: z.string(),
  execDuration: z.number(),
  execMemUse: z.number(),
  createdAt: z.string(),
  updatedAt: z.string(),
  userRefer: z.number(),
});
