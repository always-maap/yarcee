import { z } from 'zod';

export const ZUser = z.object({
  id: z.string(),
  name: z.string(),
  username: z.string(),
});
