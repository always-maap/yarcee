import { z } from 'zod';

export const ZSandbox = z.object({
  code: z.string(),
  createdAt: z.string(),
  id: z.number(),
  language: z.string(),
  name: z.string(),
  updatedAt: z.string(),
  userRefer: z.number(),
});
