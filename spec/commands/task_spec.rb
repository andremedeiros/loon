describe 'Commands' do
  describe '<task>', command: true do
    hello = {'description' => 'Says hello', 'command' => 'echo Hello Loon'}
    tasks = {'hello' => hello}

    it "should run a task from a payload" do
      with_payload(tasks: tasks) do
        loon %w(hello)
        assert_stdout "Hello Loon"
      end
    end

    it "should show the task when running loon from the project dir" do
      with_payload(tasks: tasks) do
        loon
        assert_stdout "hello"
        assert_stdout "Says hello"
      end
    end
  end
end
