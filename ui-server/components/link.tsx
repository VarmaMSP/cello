import Router from 'next/router'
import React, { cloneElement, useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { Dispatch } from 'redux'
import * as T from 'types/actions'
import { UrlObject } from 'url'

interface LinkChild {
  children: JSX.Element
}

interface LinkProps extends LinkChild {
  href: string | UrlObject
  as?: string | UrlObject
  prefetch?: boolean
  scroll?: boolean
}

export const Link: React.FC<LinkProps> = (props) => {
  const dispatch = useDispatch<Dispatch<T.AppActions>>()

  useEffect(() => {
    if (!!props.prefetch) {
      typeof props.href === 'string'
        ? Router.prefetch(props.href)
        : Router.prefetch(props.href.pathname!)
    }
  })

  return cloneElement(props.children, {
    href: typeof props.href === 'string' ? props.href : props.as,
    onClick: (e: React.SyntheticEvent<HTMLAnchorElement>) => {
      e.preventDefault()

      dispatch({
        type: T.HISTORY_PUSH_ENTRY,
        entry: {
          urlPath: Router.asPath,
          scrollY: window.scrollY,
        },
      })

      Router.push(props.href, props.as, {
        scroll: props.scroll,
      })
    },
  })
}

export const PodcastLink: React.FC<LinkChild & { podcastUrlParam: string }> = ({
  children,
  podcastUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/podcasts',
        query: { podcastUrlParam: podcastUrlParam },
      }}
      as={`/podcasts/${podcastUrlParam}`}
    >
      {children}
    </Link>
  )
}

export const EpisodeLink: React.FC<LinkChild & { episodeUrlParam: string }> = ({
  children,
  episodeUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/episodes',
        query: { episodeUrlParam: episodeUrlParam },
      }}
      as={`/episodes/${episodeUrlParam}`}
    >
      {children}
    </Link>
  )
}

export const ChartLink: React.FC<LinkChild & { chartUrlParam: string }> = ({
  children,
  chartUrlParam,
}) => {
  return (
    <Link
      href={{
        pathname: '/charts',
        query: { chartUrlParam: chartUrlParam },
      }}
      as={`/charts/${chartUrlParam}`}
    >
      {children}
    </Link>
  )
}

export const PlaylistLink: React.FC<LinkChild & {
  playlistUrlParam: string
}> = ({ children, playlistUrlParam }) => {
  return (
    <Link
      href={{
        pathname: '/playlists',
        query: { playlistUrlParam: playlistUrlParam },
      }}
      as={`/playlists/${playlistUrlParam}`}
    >
      {children}
    </Link>
  )
}
