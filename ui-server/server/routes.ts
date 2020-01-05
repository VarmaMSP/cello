import { ParameterizedContext as KoaContext } from 'koa'
import Router from 'koa-router'
import NextServer from 'next/dist/next-server/server/next-server'
import { ParsedUrlQuery } from 'querystring'

export function registerRoutes(app: NextServer, router: Router) {
  const servePage = makeServePage(app)
  const serveBuildFiles = makeServeBuildFiles(app)

  // Index page
  router.get('/', servePage('/', 'public,max-age=7200,must-revalidate'))

  // Subscriptions Feed Page
  router.get(
    '/subscriptions',
    servePage('/subscriptions', 'public,max-age=86400,must-revalidate'),
  )

  // History Feed Page
  router.get(
    '/history',
    servePage('/history', 'public,max-age=86400,must-revalidate'),
  )

  // Playlist Library Page
  router.get(
    '/playlists',
    servePage('/playlists', 'public,max-age=86400,must-revalidate'),
  )

  // Podcast Page
  router.get(
    '/podcasts/:podcastUrlParam/:activeTab*',
    servePage('/podcasts', 'public,max-age=3600,must-revalidate', (ctx) => ({
      podcastUrlParam: ctx.params['podcastUrlParam'],
      activeTab: ctx.params['activeTab'],
    })),
  )

  // Episode Page
  router.get(
    '/episodes/:episodeUrlParam/:activeTab*',
    servePage('/episodes', 'public,max-age=7200,must-revalidate', (ctx) => ({
      episodeUrlParam: ctx.params['episodeUrlParam'],
      activeTab: ctx.params['activeTab'],
    })),
  )

  // Playlist Page
  router.get(
    '/playlists/:playlistUrlParam/:activeTab*',
    servePage('/playlists', 'public,max-age=7200,must-revalidate', (ctx) => ({
      playlistUrlParam: ctx.params['playlistUrlParam'],
      activeTab: ctx.params['activeTab'],
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
      resultType: ctx.request.query['type'] || 'episode',
      sortBy: ctx.request.query['sort_by'] || 'relevance',
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
