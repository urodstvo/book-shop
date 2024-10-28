FROM node:22.9-alpine AS base

FROM base AS deps
RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY . .

RUN corepack enable pnpm 

WORKDIR /app/apps/frontend
RUN pnpm i --frozen-lockfile
RUN pnpm add sharp
RUN pnpm run build


FROM base AS runner
WORKDIR /app

ENV NODE_ENV=production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=deps --chown=nextjs:nodejs /app/apps/frontend/.next/standalone ./
COPY --from=deps --chown=nextjs:nodejs /app/apps/frontend/.next ./apps/frontend/_next
COPY --from=deps --chown=nextjs:nodejs /app/apps/frontend/.next/static ./apps/frontend/.next/static
COPY --from=deps --chown=nextjs:nodejs /app/apps/frontend/.next/static ./apps/frontend/.next/standalone/.next/static

USER nextjs

EXPOSE 3000
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"
CMD ["node", "./apps/frontend/server.js"]