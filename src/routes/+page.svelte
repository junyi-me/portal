<script lang="ts">
  type LinkType = {
    url: string;
    text: string;
  };
  const links: LinkType[] = [
    { url: 'https://profile.junyi.me/', text: 'Profile' },
    { url: 'https://blog.junyi.me/', text: 'Blog' },
  ];

  type SocialType = {
    url: string;
    icon: string;
    iconType?: string;
  };
  const socials: SocialType[] = [
    { url: "https://github.com/jywang99", icon: "github" },
    { url: "https://www.linkedin.com/in/junyi-wang-976a94199", icon: "linkedin" },
    { url: "mailto:junyi.wang.007@gmail.com", icon: "envelope", iconType: "solid" },
  ];

  const linkWidth = links.map(link => link.text.length * 0.5).reduce((a, b) => a + b, 0);
</script>

<div class="backdrop">
  <img src="/backdrop.gif" alt="backdrop" width="100%" height="100%">
  <div></div>
</div>
<div class="container">
  <div class="title">
    <h1>JUNYI WANG</h1>
    <h3>Software Engineer</h3>
  </div>
  <div class="section">
    <div class="links">
      {#each links as link}
        <a href={link.url} style="width: {linkWidth}em">{link.text}</a>
      {/each}
    </div>
  </div>
  <div class="section">
    <div class="ext-links">
      {#each socials as link}
        <a href={link.url} aria-label="link">
          <i class="{link.iconType ? `fa-${link.iconType}` : "fab"} fa-{link.icon}"></i>
        </a>
      {/each}
    </div>
  </div>

  <div class="credits">
    Backdrop by <a href="https://wallpapersafari.com/user/jtucker33/">Jessica</a> 
    on <a href="https://wallpapersafari.com/minecraft-gif-wallpapers/#oGMpdV">WallpaperSafari</a>
  </div>
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    color: var(--fg-1);
  }

  .links a {
    text-align: center;
  }

  a:link, a:visited {
    color: var(--fg-1);
  }

  .title {
    margin-bottom: 3em;
  }

  .title h1, .title h3 {
    margin: 0;
    text-align: center;
  }

  .title h3 {
    font-size: 1.5rem;
    color: var(--fg-2);
  }

  h1 {
    font-size: 3.75rem;
  }

  @media (max-width: 768px) {
    h1 {
      font-size: 2.5rem;
    }
  }

  h1 {
    text-align: center;
  }

  a {
    text-decoration: none;
  }

  .section {
    margin-bottom: 3em;
  }

  .links, .ext-links {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }

  .links {
    gap: 1em;
  }

  .ext-links {
    gap: 2em;
  }

  .links a {
    padding: 0.5em  1em;
    border: 1.5px solid var(--fg-1);
    border-radius: 2em;
    font-weight: 800;
  }

  a:hover {
    background-color: var(--accent-alpha);
    border-color: var(--accent);
    color: #ffccff;
  }

  .ext-links .link {
    padding: 6px 8px;
    transition: color 0.2s;
  }

  .ext-links .link:hover {
    color: var(--fg-2);
  }

  .ext-links i {
    font-size: 1.5rem;
  }

  .ext-links a:hover {
    color: var(--accent-light);
  }

  .backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
  }

  .backdrop img {
    object-fit: cover;
  }

  .backdrop div {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(5px);
  }

  .credits {
    position: fixed;
    bottom: 0;
    left: 0;
    padding: 1em;
    color: var(--fg-2);
  }

  .credits a {
    text-decoration: underline;
  }
</style>

