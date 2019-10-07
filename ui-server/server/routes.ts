import { ParameterizedContext as KoaContext } from 'koa'
import Router from 'koa-router'
import NextServer from 'next-server/dist/server/next-server'
import { ParsedUrlQuery } from 'querystring'

export function registerRoutes(app: NextServer, router: Router) {
  const servePage = makeServePage(app)
  const serveBuildFiles = makeServeBuildFiles(app)

  // Index page
  router.get('/', servePage('/', 'public,max-age=7200,must-revalidate'))

  // Feed Page
  router.get(
    '/feed',
    servePage('/feed', 'public,max-age=86400,must-revalidate'),
  )

  // Subscriptions Page
  router.get(
    '/subscriptions',
    servePage('/subscriptions', 'public,max-age=86400,must-revalidate'),
  )

  // Podcast Page
  router.get(
    '/podcasts/:podcastId',
    servePage('/podcasts', 'public,max-age=3600,must-revalidate', (ctx) => ({
      podcastId: ctx.params['podcastId'],
    })),
  )

  // Results Page
  router.get(
    '/results',
    servePage('/results', 'public,max-age=300,must-revalidate', (ctx) => ({
      query: ctx.request.query['query'],
    })),
  )

  router.get('*', serveBuildFiles)
}

function makeServePage(app: NextServer) {
  return (
    page: string,
    cacheControl: string,
    query?: (ctx: KoaContext) => ParsedUrlQuery,
  ) => {
    return async (ctx: KoaContext) => {
      if (process.env.NODE_ENV === 'production') {
        ctx.set({ 'Cache-Control': cacheControl })
      }

      await app.render(ctx.req, ctx.res, page, (query && query(ctx)) || {})
      ctx.respond = false
    }
  }
}

function makeServeBuildFiles(app: NextServer) {
  return async (ctx: KoaContext) => {
    ctx.set({ 'Cache-Control': 'public,max-age=31536000,immutable' })
    await app.getRequestHandler()(ctx.req, ctx.res)
    ctx.respond = false
  }
}
