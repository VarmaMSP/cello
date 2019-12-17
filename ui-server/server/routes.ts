import { ParameterizedContext as KoaContext } from 'koa'
import Router from 'koa-router'
import NextServer from 'next/dist/next-server/server/next-server'
import { ParsedUrlQuery } from 'querystring'

export function registerRoutes(app: NextServer, router: Router) {
  const servePage = makeServePage(app)
  const serveBuildFiles = makeServeBuildFiles(app)

  // Index page
  router.get('/', servePage('/', 'public,max-age=7200,must-revalidate'))

  // Subscriptions Page
  router.get(
    '/subscriptions',
    servePage('/subscriptions', 'public,max-age=86400,must-revalidate'),
  )

  // Podcast Page
  router.get(
    '/podcasts/:podcastUrlParam/:activeTab*',
    servePage('/podcasts', 'public,max-age=3600,must-revalidate', (ctx) => ({
      podcastUrlParam: ctx.params['podcastUrlParam'],
      activeTab: ctx.params['activeTab'] || 'about',
    })),
  )

  // Episode Page
  router.get(
    '/episodes/:episodeUrlParam',
    servePage('/episodes', 'public,max-age=7200,must-revalidate', (ctx) => ({
      episodeUrlParam: ctx.params['episodeurlParam'],
    })),
  )

  // Charts Page
  router.get(
    '/charts/:chartId',
    servePage('/charts', 'public,max-age=1200,must-revalidate', (ctx) => ({
      chartId: ctx.params['chartId'],
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
