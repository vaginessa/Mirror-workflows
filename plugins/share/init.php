<?php
class Share extends Plugin {
	private $host;

	function about() {
		return array(1.0,
			"Share article by unique URL",
			"fox");
	}

	/* @var PluginHost $host */
	function init($host) {
		$this->host = $host;

		$host->add_hook($host::HOOK_ARTICLE_BUTTON, $this);
		$host->add_hook($host::HOOK_PREFS_TAB_SECTION, $this);
	}

	function get_js() {
		return file_get_contents(__DIR__ . "/share.js");
	}

	function get_css() {
		return file_get_contents(__DIR__ . "/share.css");
	}

	function get_prefs_js() {
		return file_get_contents(__DIR__ . "/share_prefs.js");
	}

	function unshare() {
		$id = $_REQUEST['id'];

		$sth = $this->pdo->prepare("UPDATE ttrss_user_entries SET uuid = '' WHERE int_id = ?
			AND owner_uid = ?");
		$sth->execute([$id, $_SESSION['uid']]);

		print __("Article unshared");
	}

	function hook_prefs_tab_section($id) {
		if ($id == "prefFeedsPublishedGenerated") {
			?>
			<hr/>

			<h2><?= __("You can disable all articles shared by unique URLs here.") ?></h2>

			<button class='alt-danger' dojoType='dijit.form.Button' onclick="return Plugins.Share.clearKeys()">
				<?= __('Unshare all articles') ?></button>
			<?php
		}
	}

	function clearArticleKeys() {
		$sth = $this->pdo->prepare("UPDATE ttrss_user_entries SET uuid = '' WHERE
			owner_uid = ?");
		$sth->execute([$_SESSION['uid']]);

		print __("Shared URLs cleared.");
	}

	function newkey() {
		$id = $_REQUEST['id'];
		$uuid = uniqid_short();

		$sth = $this->pdo->prepare("UPDATE ttrss_user_entries SET uuid = ? WHERE int_id = ?
			AND owner_uid = ?");
		$sth->execute([$uuid, $id, $_SESSION['uid']]);

		print json_encode(["link" => $uuid]);
	}

	function hook_article_button($line) {
		$icon_class = !empty($line['uuid']) ? "is-shared" : "";

		return "<i class='material-icons icon-share share-icon-".$line['int_id']." $icon_class'
			style='cursor : pointer' onclick=\"Plugins.Share.shareArticle(".$line['int_id'].")\"
			title='".__('Share by URL')."'>link</i>";
	}

	function shareDialog() {
		$id = (int)clean($_REQUEST['id'] ?? 0);

		$sth = $this->pdo->prepare("SELECT uuid FROM ttrss_user_entries WHERE int_id = ?
			AND owner_uid = ?");
		$sth->execute([$id, $_SESSION['uid']]);

		if ($row = $sth->fetch()) {
			$uuid = $row['uuid'];

			if (!$uuid) {
				$uuid = uniqid_short();

				$sth = $this->pdo->prepare("UPDATE ttrss_user_entries SET uuid = ? WHERE int_id = ?
					AND owner_uid = ?");
				$sth->execute([$uuid, $id, $_SESSION['uid']]);
			}

			$url_path = get_self_url_prefix() . "/public.php?op=share&key=$uuid";

			?>

			<header><?= __("You can share this article by the following unique URL:") ?></header>

			<section>
				<div class='panel text-center'>
					<a class='target-url' href="<?= htmlspecialchars($url_path) ?>"
						target='_blank' rel='noopener noreferrer'><?= htmlspecialchars($url_path) ?></a>
				</div>
			</section>

			<?php

		} else {
			print format_error(__("Article not found."));
		}

		?>
		<footer class='text-center'>
			<?= \Controls\button_tag(__('Unshare article'), '', ['class' => 'alt-danger', 'onclick' => "App.dialogOf(this).unshare()"]) ?>
			<?= \Controls\button_tag(__('Generate new URL'), '', ['onclick' => "App.dialogOf(this).newurl()"]) ?>
			<?= \Controls\submit_tag(__("Close this window")) ?>
		</footer>
		<?php
	}

	function api_version() {
		return 2;
	}

}
