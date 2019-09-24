import Koa from 'koa'
import Router from 'koa-router'
import next from 'next'

const app = next({ dev: process.env.NODE_ENV !== 'production' })
const handle = app.getRequestHandler()
const router = new Router()

router.get('/', async (ctx) => {
  ctx.set({
    'Cache-Control': '',
  })
  await app.render(ctx.req, ctx.res, '/')
  ctx.respond = false
})

router.get('/podcasts/:podcastId', async (ctx) => {
  await app.render(ctx.req, ctx.res, '/podcasts', <any>{
    podcastId: ctx.params['podcastId'],
  })
  ctx.respond = false
})

router.get('/results', async (ctx) => {
  await app.render(ctx.req, ctx.res, '/results', {
    query: ctx.request.query['query'],
  })
  ctx.respond = false
})

router.get('*', async (ctx) => {
  await handle(ctx.req, ctx.res)
  ctx.respond = false
})

const server = new Koa()

server
  .use(async (ctx, next) => {
    ctx.res.statusCode = 200
    await next()
  })
  .use(router.routes())
  .use(router.allowedMethods())

app.prepare().then(() => {
  const port = parseInt(process.env.PORT || '8082', 10)
  server.listen(port, () => {
    console.log(`UI server running on port ${port}`)
  })
})
