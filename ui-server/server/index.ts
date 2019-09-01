import Koa from 'koa'
import Router from 'koa-router'
import next from 'next'
import bodyParser from 'koa-bodyparser'
import { ParsedUrlQuery } from 'querystring'

const app = next({ dev: process.env.NODE_ENV !== 'production' })
const handle = app.getRequestHandler()
const router = new Router()

router.get('/podcasts', async (ctx) => {
  const body: ParsedUrlQuery = ctx.request.body as any
  await app.render(ctx.req, ctx.res, '/podcast', body)
  ctx.respond = false
})

router.get('*', async (ctx) => {
  await handle(ctx.req, ctx.res)
  ctx.respond = false
})

const server = new Koa()

server.use(bodyParser())
server.use(router.routes())
server.use(router.allowedMethods())

app.prepare().then(() => {
  const port = parseInt(process.env.PORT || '8081', 10)
  server.listen(port, () => {
    console.log('UI server running on port 8081')
  })
})
