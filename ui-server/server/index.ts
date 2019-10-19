import Koa from 'koa'
import favicon from 'koa-favicon'
import Router from 'koa-router'
import next from 'next'
import path from 'path'
import { registerRoutes } from './routes'

const app = next({ dev: process.env.NODE_ENV !== 'production' })
const server = new Koa()
const router = new Router()
const staticDir = path.resolve('__dirname', '..', 'static')

registerRoutes(app, router)

server
  .use(async (ctx, next) => {
    ctx.res.statusCode = 200
    await next()
  })
  .use(favicon(path.resolve(staticDir, 'favicon.ico')))
  .use(router.routes())
  .use(router.allowedMethods())

app.prepare().then(() => {
  const port = +(process.env.PORT || '8082')
  server.listen(port, () => {
    console.log(`UI server running on port ${port}`)
  })
})
