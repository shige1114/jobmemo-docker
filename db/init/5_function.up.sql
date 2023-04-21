CREATE OR REPLACE FUNCTION set_selections_recruits_fail()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.fail = TRUE THEN -- テーブルAのカラムcと、テーブルBのカラムkが異なる場合にのみ、テーブルAのカラムcに、テーブルBのカラムkの値をセットする
    UPDATE recruits SET reject = NEW.fail WHERE recruits.users_id = NEW.users_id AND recruits.companies_id = NEW.companies_id;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER link_selections_recruits
AFTER INSERT ON selections
FOR EACH ROW
EXECUTE FUNCTION set_selections_recruits_fail();

