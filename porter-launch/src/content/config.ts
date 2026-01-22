import { defineCollection, z } from "astro:content";

const docsCollection = defineCollection({
  type: "content",
  schema: z.object({
    title: z.string(),
    description: z.string(),
    category: z.enum([
      "getting-started",
      "concepts",
      "configuration",
      "api",
      "examples",
      "guides"
    ]),
    order: z.number().optional(),
    updated: z.date().optional()
  })
});

const marketingCollection = defineCollection({
  type: "data",
  schema: z.object({
    hero: z
      .object({
        headline: z.string(),
        subhead: z.string(),
        cta: z.string()
      })
      .optional(),
    features: z
      .array(
        z.object({
          title: z.string(),
          description: z.string(),
          icon: z.string()
        })
      )
      .optional(),
    useCases: z
      .array(
        z.object({
          title: z.string(),
          description: z.string(),
          example: z.string()
        })
      )
      .optional()
  })
});

export const collections = {
  docs: docsCollection,
  marketing: marketingCollection
};
