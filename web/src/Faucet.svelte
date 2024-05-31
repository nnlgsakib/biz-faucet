<script>
    import { onMount } from 'svelte';
    import { getAddress } from '@ethersproject/address';
    import { formatEther } from '@ethersproject/units';
    import { setDefaults as setToast, toast } from 'bulma-toast';

    // Define an array of background images
    const backgrounds = [
      '/1.webp',
      '/2.webp',
      '/3.webp'
      // Add more background images here
    ];

    let distribution = 10;
    let address = null;
    let faucetInfo = {
      account: '0x0000000000000000000000000000000000000000',
      network: 'testnet',
      payout: 1000
    };

    $: document.title = `BIZ ${capitalize(faucetInfo.network)} Faucet`;

    // Function to select a random background imaged:\contents\ass\3.gif
    function getRandomBackground() {
      return backgrounds[Math.floor(Math.random() * backgrounds.length)];
    }

    let backgroundImage = getRandomBackground(); // Initial random background image

    onMount(async () => {
      const res = await fetch('/api/info');
      faucetInfo = await res.json();
      faucetInfo.payout = parseInt(formatEther(faucetInfo.payout));
    });

    setToast({
      position: 'bottom-center',
      dismissible: true,
      pauseOnHover: true,
      closeOnClick: false,
      animate: { in: 'fadeIn', out: 'fadeOut' },
    });

    async function handleRequest() {
      try {
        address = getAddress(address);
      } catch (error) {
        toast({ message: error.reason, type: 'is-warning' });
        return;
      }

      let formData = new FormData();
      formData.append('address', address);
      const res = await fetch('/api/claim', {
        method: 'POST',
        body: formData,
      });
      let message = await res.text();
      let type = res.ok ? 'is-success' : 'is-warning';
      toast({ message, type });
    }

    function capitalize(str) {
      const lower = str.toLowerCase();
      return str.charAt(0).toUpperCase() + lower.slice(1);
    }
</script>

<main>
    <section class="hero is-info is-fullheight" style="background: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.8)), url({backgroundImage}) center center / cover no-repeat fixed;">
      <div class="hero-head">
        <nav class="navbar">
          <div class="container">
            <div class="navbar-brand">
              <a class="navbar-item" href=".">
                <span class="icon">
                  <i class="fa fa-bath" />
                </span>
                <span><b>BIZ (testnet) Faucet</b></span>
              </a>
            </div>
            <div id="navbarMenu" class="navbar-menu">
              <div class="navbar-end">
                <span class="navbar-item">
                  <a
                    class="button is-white is-outlined"
                    href="https://github.com/BIZchain-labs/biz-faucet"
                  >
                    <span class="icon">
                      <i class="fa fa-github" />
                    </span>
                    <span>View Source</span>
                  </a>
                </span>
              </div>
            </div>
          </div>
        </nav>
      </div>

      <div class="hero-body">
        <div class="container has-text-centered">
          <div class="column is-6 is-offset-3">
            <h1 class="title is-spaced">
              <span style="color: #00d1b2;">Receive</span> {distribution} tBIZ <span style="color: #00d1b2;">per request</span>
            </h1>
            <h2 class="subtitle">
              <span style="color: #00d1b2;">Serving from</span> {faucetInfo.account}
            </h2>
            <div class="box animated fadeInUp">
              <div class="field is-grouped">
                <p class="control is-expanded">
                  <input
                    bind:value={address}
                    class="input is-rounded is-black"
                    type="text"
                    placeholder="Enter your address"
                  />
                </p>
                <p class="control">
                  <button
                    on:click={handleRequest}
                    class="button is-primary is-rounded"
                  >
                    Request
                  </button>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
</main>

<style>
    main {
      overflow: hidden; /* Hide scrollbars */
    }
    .hero.is-info {
      background-size: cover;
    }
    .hero .subtitle {
      padding: 3rem 0;
      line-height: 1.5;
    }
    .box {
      border-radius: 19px;
      animation-duration: 1s;
      animation-delay: 0.5s;
      background-color: rgba(255, 255, 255, 0.1);
      border: 2px solid #00d1b2;
      box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.4);
    }
    .animated {
      animation-fill-mode: both;
    }
    .fadeInUp {
      animation-name: fadeInUp;
    }
    @keyframes fadeInUp {
      from {
        opacity: 0;
        transform: translateY(20px);
      }
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }
</style>
