<script lang="ts">
  import '@fortawesome/fontawesome-free/css/all.min.css'

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
        <div class="link">
          <a href={link.url}>{link.text}</a>
        </div>
      {/each}
    </div>
  </div>
  <div class="section">
    <div class="ext-links">
      {#each socials as link}
        <div class="link">
          <a href={link.url} aria-label="link">
            <i class="{link.iconType ? `fa-${link.iconType}` : "fab"} fa-{link.icon}"></i>
          </a>
        </div>
      {/each}
    </div>
  </div>
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    color: #fff;
  }

  a:link, a:visited {
    color: #fff;
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
    color: #BEBEBE;
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
    gap: 10px;
  }

  .links .link {
    display: flex;
    align-items: center;
    justify-content: center;
    background: rgba(0, 0, 0, 0.5);
    width: 6em;
    padding: 6px 8px;
    border: 1px solid #fff;
    border-radius: 20px;
  }

  .ext-links .link {
    padding: 6px 8px;
  }

  .ext-links .link i {
    font-size: 1.5rem;
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
</style>
